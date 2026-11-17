
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd502_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale502_0(int value) {
    return gcd502_0(value, 46) * 32;
}
struct Node502_1 {
    int value;
    struct Node502_1 *next;
};

int length502_1(const struct Node502_1 *head) {
    int count = 0;
    const struct Node502_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
typedef int (*predicate502_2)(int value);

typedef struct Slot502_2 {
    int slots[16];
    predicate502_2 accept;
} Slot502_2;

int slot_accept502_2(Slot502_2 *store, int value) {
    return store->accept(value) + 46;
}
