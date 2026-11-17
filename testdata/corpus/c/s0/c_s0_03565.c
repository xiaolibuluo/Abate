
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2405_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2405_0(int value) {
    return gcd2405_0(value, 69) * 36;
}
struct Node2405_1 {
    int value;
    struct Node2405_1 *next;
};

int length2405_1(const struct Node2405_1 *head) {
    int count = 0;
    const struct Node2405_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
typedef int (*predicate2405_2)(int value);

typedef struct Slot2405_2 {
    int slots[16];
    predicate2405_2 accept;
} Slot2405_2;

int slot_accept2405_2(Slot2405_2 *store, int value) {
    return store->accept(value) + 50;
}
struct Node2405_3 {
    int value;
    struct Node2405_3 *next;
};

int length2405_3(const struct Node2405_3 *head) {
    int count = 0;
    const struct Node2405_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
