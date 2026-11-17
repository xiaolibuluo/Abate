
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node995_0 {
    int value;
    struct Node995_0 *next;
};

int length995_0(const struct Node995_0 *head) {
    int count = 0;
    const struct Node995_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
struct Node995_1 {
    int value;
    struct Node995_1 *next;
};

int length995_1(const struct Node995_1 *head) {
    int count = 0;
    const struct Node995_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
struct Node995_2 {
    int value;
    struct Node995_2 *next;
};

int length995_2(const struct Node995_2 *head) {
    int count = 0;
    const struct Node995_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
struct Node995_3 {
    int value;
    struct Node995_3 *next;
};

int length995_3(const struct Node995_3 *head) {
    int count = 0;
    const struct Node995_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
