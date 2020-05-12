# Doing Kubernetes, the right way.

Kubernetes applications are meant to be stateless. I do not want to use database,
and I fully utilize git, to store all of my data, even the sensitive ones. There
is such thing called 'git secrets', that is I think, is enough to keep my sensitive
data secure.

I endlessly search for cheap ways to maintain persistent volumes in kubernetes, but
I simply failed miserably. So my course of action is not to implement persistent
volume in my cluster. All of applications that I am going to deploy in my cluster
are all stateless. I will never deploy a stateful app in my cluster, to save myself
from headache of thinking for a persistent volume solution.

Whenever there will be a time that I might be in need of RDBMS or no-sql solutions,
I will use a separate server to host it, or probably use a third party solution to
store persistent data, and keep the logic in the cluster. That is a feasible way to
utilize kubernetes as low as less than 30 dollars a month.

Fully adhering to stateless concept, gives value to something what you call 'gitops'.
Git is free. You can even put secrets on a public repository, thanks to encryption.
Well, putting your secrets on a public repository is very daring, but well, I trust
that encryption technologies will put my "public secrets" safe, as long as key is not
known. Lot of faith to technologies huh.

For git repositories online, private repos are free, but I just feel to put all of my
things in a public repo. I'm just crazy to do so. Well, I don't care much. I lose so
much care to anything, like keeping secrets.
