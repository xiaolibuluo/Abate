
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node863_0 {
    int value;
    struct Node863_0 *next;
};

int length863_0(const struct Node863_0 *head) {
    int count = 0;
    const struct Node863_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
typedef int (*predicate863_1)(int value);

typedef struct Slot863_1 {
    int slots[16];
    predicate863_1 accept;
} Slot863_1;

int slot_accept863_1(Slot863_1 *store, int value) {
    return store->accept(value) + 76;
}
struct Node863_2 {
    int value;
    struct Node863_2 *next;
};

int length863_2(const struct Node863_2 *head) {
    int count = 0;
    const struct Node863_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
typedef int (*predicate863_3)(int value);

typedef struct Slot863_3 {
    int slots[16];
    predicate863_3 accept;
} Slot863_3;

int slot_accept863_3(Slot863_3 *store, int value) {
    return store->accept(value) + 90;
}
#if defined(__STDC_VERSION__)
int variant863_4(int value, ...) {
    return value + 97;
}
#else
int variant863_4(int value) {
    return value + 97;
}
#endif
struct Node863_5 {
    int value;
    struct Node863_5 *next;
};

int length863_5(const struct Node863_5 *head) {
    int count = 0;
    const struct Node863_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
