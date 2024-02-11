<template>
    <div class="profile-actions">
      <!-- Update User Details with Modal -->
      <div>
        <h2>Update User Details</h2>
        <button class="profile-button" @click="openModal">Update Details</button>
  
        <!-- Modal -->
        <div v-if="isModalOpen" class="modal">
          <div class="modal-content">
            <span class="close" @click="closeModal">&times;</span>
            <form @submit.prevent="updateUserDetails">
              <label class="profile-label" for="modalusername">Username:</label>
              <input class="profile-input" v-model="modalusername" type="text" id="modalusername" />
  
              <label class="profile-label" for="modalfirstname">First Name:</label>
              <input class="profile-input" v-model="modalfirstname" type="text" id="modalfirstname" />
  
              <label class="profile-label" for="modallastname">Last Name:</label>
              <input class="profile-input" v-model="modallastname" type="text" id="modallastname" />
  
              <label class="profile-label" for="modalemail">Email:</label>
              <input class="profile-input" v-model="modalemail" type="email" id="modalemail" />
  
              <button class="profile-button" type="submit">Submit</button>
            </form>
          </div>
        </div>
      </div>
  
      <!-- Upload Image -->
      <div>
        <h2>Upload Image</h2>
        <form @submit.prevent="uploadImage">
          <label class="profile-label" for="image">Choose Image:</label>
          <input class="profile-input" type="file" @change="handleFileChange" accept="image/*" id="image" required />
  
          <button class="profile-button" type="submit">Upload Image</button>
        </form>
      </div>
  
      <!-- Follow Box -->
      <div>
        <h2>Follow Box</h2>
        <label class="profile-label" for="followInput">Follow:</label>
        <input class="profile-input" v-model="followInput" type="text" id="followInput" />
        <button class="profile-button" @click="followUser">Follow</button>
        <button class="profile-button" @click="unfollowUser">Unfollow</button>
      </div>
  
      <!-- Ban Box -->
      <div>
        <h2>Ban Box</h2>
        <label class="profile-label" for="banInput">Ban:</label>
        <input class="profile-input" v-model="banInput" type="text" id="banInput" />
        <button class="profile-button" @click="banUser">Ban</button>
        <button class="profile-button" @click="unbanUser">Unban</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        followInput: "", // Input for following a user
        banInput: "", // Input for banning a user
        isModalOpen: false,
        modalusername: "",
        modalfirstname: "",
        modallastname: "",
        modalemail: "",
      };
    },
    methods: {
      openModal() {
        this.isModalOpen = true;
      },
  
      closeModal() {
        this.isModalOpen = false;
      },
  
      async updateUserDetails() {
        const userID = this.$route.params.userID;
      const updatedFields = {};

      if (this.modalusername) updatedFields.username = this.modalusername;
      if (this.modalfirstname) updatedFields.firstname = this.modalfirstname;
      if (this.modallastname) updatedFields.lastname = this.modallastname;
      if (this.modalemail) updatedFields.email = this.modalemail;

      try {
        await fetch(`http://localhost:3000/users/${userID}`, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(updatedFields),
        });

        alert("User details updated successfully!");
        this.closeModal();
        this.fetchUserProfile();
      } catch (error) {
        console.error("Error updating user details", error);
      }
      },
  
      handleFileChange(event) {
        this.selectedImage = event.target.files[0];
      },
  
      async uploadImage() {
        const userID = this.$route.params.userID;
        const formData = new FormData();
        formData.append("file", this.selectedImage);
  
        try {
          await fetch(`http://localhost:3000/users/${userID}/image`, {
            method: 'POST',
            body: formData,
          });
          alert("Image uploaded successfully!");
          // Optionally, you can fetch the updated user profile after uploading the image
          this.fetchUserProfile();
        } catch (error) {
          console.error("Error uploading image", error);
        }
      },
  
      async followUser() {
        const userID = this.$route.params.userID;
      try {
        await fetch(`http://localhost:3000/users/${userID}/follow?FollowListOwnerID=${this.followInput}`, {
          method: 'POST',
        });
        alert(`You are now following user with ID ${this.followInput}!`);
        this.fetchUserProfile(); // Optionally, you can fetch the updated user profile
      } catch (error) {
        console.error("Error following user", error);
      }
      },
  
      async unfollowUser() {
        const userID = this.$route.params.userID;
      try {
        await fetch(`http://localhost:3000/users/${userID}/follow?FollowListOwnerID=${this.followInput}`, {
          method: 'DELETE',
        });
        alert(`You have unfollowed user with ID ${this.followInput}!`);
        this.fetchUserProfile(); // Optionally, you can fetch the updated user profile
      } catch (error) {
        console.error("Error unfollowing user", error);
      }
      },
  
      async banUser() {
        const userID = this.$route.params.userID;
      try {
        await fetch(`http://localhost:3000/users/${userID}/ban?BannedUserID=${this.banInput}`, {
          method: 'POST',
        });
        alert(`You have banned user with ID ${this.banInput}!`);
        this.fetchUserProfile(); // Optionally, you can fetch the updated user profile
      } catch (error) {
        console.error("Error banning user", error);
      }
      },
  
      async unbanUser() {
        const userID = this.$route.params.userID;
      try {
        await fetch(`http://localhost:3000/users/${userID}/ban?BannedUserID=${this.banInput}`, {
          method: 'DELETE',
        });
        alert(`You have unbanned user with ID ${this.banInput}!`);
        this.fetchUserProfile(); // Optionally, you can fetch the updated user profile
      } catch (error) {
        console.error("Error unbanning user", error);
      }

      },
    },
  };
  </script>
  
  <style scoped>
  /* Styles for profile actions */
  .profile-actions {
    display: flex;
    justify-content: space-around;
    margin: 20px;
  }
  
  /* Styles for labels, inputs, and buttons */
  .profile-label {
    display: block;
    margin-bottom: 5px;
  }
  
  .profile-input {
    width: 100%;
    padding: 8px;
    margin-bottom: 15px;
  }
  
  .profile-button {
    background-color: #4267b2;
    color: #fff;
    padding: 10px;
    border: none;
    cursor: pointer;
  }
  </style>
  