
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2166_0 {
    int value;
    struct Node2166_0 *next;
};

int length2166_0(const struct Node2166_0 *head) {
    int count = 0;
    const struct Node2166_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
int gcd2166_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2166_1(int value) {
    return gcd2166_1(value, 54) * 40;
}
struct Node2166_2 {
    int value;
    struct Node2166_2 *next;
};

int length2166_2(const struct Node2166_2 *head) {
    int count = 0;
    const struct Node2166_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
int gcd2166_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2166_3(int value) {
    return gcd2166_3(value, 15) * 54;
}
typedef int (*predicate2166_4)(int value);

typedef struct Slot2166_4 {
    int slots[16];
    predicate2166_4 accept;
} Slot2166_4;

int slot_accept2166_4(Slot2166_4 *store, int value) {
    return store->accept(value) + 61;
}
