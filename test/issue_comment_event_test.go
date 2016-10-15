package github

import (
	"testing"
	"encoding/json"
)

func TestIssueCommentEvent(t *testing.T) {
	event := `{
	  "action": "created",
	  "issue": {
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
	    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/labels{/name}",
	    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/comments",
	    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2/events",
	    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2",
	    "id": 73464126,
	    "number": 2,
	    "title": "Spelling error in the README file",
	    "user": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "labels": [
	      {
	        "url": "https://api.github.com/repos/baxterthehacker/public-repo/labels/bug",
	        "name": "bug",
	        "color": "fc2929"
	      }
	    ],
	    "state": "open",
	    "locked": false,
	    "assignee": null,
	    "milestone": null,
	    "comments": 1,
	    "created_at": "2015-05-05T23:40:28Z",
	    "updated_at": "2015-05-05T23:40:28Z",
	    "closed_at": null,
	    "body": "It looks like you accidently spelled 'commit' with two 't's."
	  },
	  "comment": {
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments/99262140",
	    "html_url": "https://github.com/baxterthehacker/public-repo/issues/2#issuecomment-99262140",
	    "issue_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/2",
	    "id": 99262140,
	    "user": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "created_at": "2015-05-05T23:40:28Z",
	    "updated_at": "2015-05-05T23:40:28Z",
	    "body": "You are totally right! I'll get this fixed right away."
	  },
	  "repository": {
	    "id": 35129377,
	    "name": "public-repo",
	    "full_name": "baxterthehacker/public-repo",
	    "owner": {
	      "login": "baxterthehacker",
	      "id": 6752317,
	      "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	      "gravatar_id": "",
	      "url": "https://api.github.com/users/baxterthehacker",
	      "html_url": "https://github.com/baxterthehacker",
	      "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	      "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	      "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	      "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	      "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	      "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	      "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	      "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	      "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	      "type": "User",
	      "site_admin": false
	    },
	    "private": false,
	    "html_url": "https://github.com/baxterthehacker/public-repo",
	    "description": "",
	    "fork": false,
	    "url": "https://api.github.com/repos/baxterthehacker/public-repo",
	    "forks_url": "https://api.github.com/repos/baxterthehacker/public-repo/forks",
	    "keys_url": "https://api.github.com/repos/baxterthehacker/public-repo/keys{/key_id}",
	    "collaborators_url": "https://api.github.com/repos/baxterthehacker/public-repo/collaborators{/collaborator}",
	    "teams_url": "https://api.github.com/repos/baxterthehacker/public-repo/teams",
	    "hooks_url": "https://api.github.com/repos/baxterthehacker/public-repo/hooks",
	    "issue_events_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/events{/number}",
	    "events_url": "https://api.github.com/repos/baxterthehacker/public-repo/events",
	    "assignees_url": "https://api.github.com/repos/baxterthehacker/public-repo/assignees{/user}",
	    "branches_url": "https://api.github.com/repos/baxterthehacker/public-repo/branches{/branch}",
	    "tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/tags",
	    "blobs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/blobs{/sha}",
	    "git_tags_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/tags{/sha}",
	    "git_refs_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/refs{/sha}",
	    "trees_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/trees{/sha}",
	    "statuses_url": "https://api.github.com/repos/baxterthehacker/public-repo/statuses/{sha}",
	    "languages_url": "https://api.github.com/repos/baxterthehacker/public-repo/languages",
	    "stargazers_url": "https://api.github.com/repos/baxterthehacker/public-repo/stargazers",
	    "contributors_url": "https://api.github.com/repos/baxterthehacker/public-repo/contributors",
	    "subscribers_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscribers",
	    "subscription_url": "https://api.github.com/repos/baxterthehacker/public-repo/subscription",
	    "commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/commits{/sha}",
	    "git_commits_url": "https://api.github.com/repos/baxterthehacker/public-repo/git/commits{/sha}",
	    "comments_url": "https://api.github.com/repos/baxterthehacker/public-repo/comments{/number}",
	    "issue_comment_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments{/number}",
	    "contents_url": "https://api.github.com/repos/baxterthehacker/public-repo/contents/{+path}",
	    "compare_url": "https://api.github.com/repos/baxterthehacker/public-repo/compare/{base}...{head}",
	    "merges_url": "https://api.github.com/repos/baxterthehacker/public-repo/merges",
	    "archive_url": "https://api.github.com/repos/baxterthehacker/public-repo/{archive_format}{/ref}",
	    "downloads_url": "https://api.github.com/repos/baxterthehacker/public-repo/downloads",
	    "issues_url": "https://api.github.com/repos/baxterthehacker/public-repo/issues{/number}",
	    "pulls_url": "https://api.github.com/repos/baxterthehacker/public-repo/pulls{/number}",
	    "milestones_url": "https://api.github.com/repos/baxterthehacker/public-repo/milestones{/number}",
	    "notifications_url": "https://api.github.com/repos/baxterthehacker/public-repo/notifications{?since,all,participating}",
	    "labels_url": "https://api.github.com/repos/baxterthehacker/public-repo/labels{/name}",
	    "releases_url": "https://api.github.com/repos/baxterthehacker/public-repo/releases{/id}",
	    "created_at": "2015-05-05T23:40:12Z",
	    "updated_at": "2015-05-05T23:40:12Z",
	    "pushed_at": "2015-05-05T23:40:27Z",
	    "git_url": "git://github.com/baxterthehacker/public-repo.git",
	    "ssh_url": "git@github.com:baxterthehacker/public-repo.git",
	    "clone_url": "https://github.com/baxterthehacker/public-repo.git",
	    "svn_url": "https://github.com/baxterthehacker/public-repo",
	    "homepage": null,
	    "size": 0,
	    "stargazers_count": 0,
	    "watchers_count": 0,
	    "language": null,
	    "has_issues": true,
	    "has_downloads": true,
	    "has_wiki": true,
	    "has_pages": true,
	    "forks_count": 0,
	    "mirror_url": null,
	    "open_issues_count": 2,
	    "forks": 0,
	    "open_issues": 2,
	    "watchers": 0,
	    "default_branch": "master"
	  },
	  "sender": {
	    "login": "baxterthehacker",
	    "id": 6752317,
	    "avatar_url": "https://avatars.githubusercontent.com/u/6752317?v=3",
	    "gravatar_id": "",
	    "url": "https://api.github.com/users/baxterthehacker",
	    "html_url": "https://github.com/baxterthehacker",
	    "followers_url": "https://api.github.com/users/baxterthehacker/followers",
	    "following_url": "https://api.github.com/users/baxterthehacker/following{/other_user}",
	    "gists_url": "https://api.github.com/users/baxterthehacker/gists{/gist_id}",
	    "starred_url": "https://api.github.com/users/baxterthehacker/starred{/owner}{/repo}",
	    "subscriptions_url": "https://api.github.com/users/baxterthehacker/subscriptions",
	    "organizations_url": "https://api.github.com/users/baxterthehacker/orgs",
	    "repos_url": "https://api.github.com/users/baxterthehacker/repos",
	    "events_url": "https://api.github.com/users/baxterthehacker/events{/privacy}",
	    "received_events_url": "https://api.github.com/users/baxterthehacker/received_events",
	    "type": "User",
	    "site_admin": false
	  }
	}`
	var issueCommentEvent IssueCommentEvent
	err := json.Unmarshal([]byte(event), &issueCommentEvent)
	if err != nil {
		t.Error(err)
	} else {
		if issueCommentEvent.Action != "created" {
			t.Error("issueCommentEvent.Action != created")
		}
		if issueCommentEvent.Issue.Title != "Spelling error in the README file" {
			t.Error("issueCommentEvent.Issue.Title != Spelling error in the README file")
		}
		if len(issueCommentEvent.Issue.Labels) != 1 {
			t.Error("len(issueCommentEvent.Issue.Labels) != 1")
		}
		if issueCommentEvent.Issue.Labels[0].Color != "fc2929" {
			t.Error("issueCommentEvent.Issue.Labels[0].Color != fc2929")
		}
		if issueCommentEvent.Comment.Url != "https://api.github.com/repos/baxterthehacker/public-repo/issues/comments/99262140" {
			t.Error("issueCommentEvent.Comment.Url != https://api.github.com/repos/baxterthehacker/public-repo/issues/comments/99262140")
		}
		repo := issueCommentEvent.Repository
		if repo.Id != 35129377 {
			t.Error("issueCommentEvent.Repository.Id != 35129377")
		}
		if repo.Owner.Login != "baxterthehacker" {
			t.Error("issueCommentEvent.Repository.Owner.Login != baxterthehacker")
		}
		sender := issueCommentEvent.Sender
		if sender.Login != "baxterthehacker" {
			t.Error("issueCommentEvent.Sender.Login != baxterthehacker")
		}
	}
}