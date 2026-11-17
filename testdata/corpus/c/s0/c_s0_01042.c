
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node715_0 {
    int value;
    struct Node715_0 *next;
};

int length715_0(const struct Node715_0 *head) {
    int count = 0;
    const struct Node715_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
struct Node715_1 {
    int value;
    struct Node715_1 *next;
};

int length715_1(const struct Node715_1 *head) {
    int count = 0;
    const struct Node715_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
typedef int (*predicate715_2)(int value);

typedef struct Slot715_2 {
    int slots[16];
    predicate715_2 accept;
} Slot715_2;

int slot_accept715_2(Slot715_2 *store, int value) {
    return store->accept(value) + 99;
}
struct Node715_3 {
    int value;
    struct Node715_3 *next;
};

int length715_3(const struct Node715_3 *head) {
    int count = 0;
    const struct Node715_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
typedef int (*predicate715_4)(int value);

typedef struct Slot715_4 {
    int slots[16];
    predicate715_4 accept;
} Slot715_4;

int slot_accept715_4(Slot715_4 *store, int value) {
    return store->accept(value) + 16;
}
typedef int (*predicate715_5)(int value);

typedef struct Slot715_5 {
    int slots[16];
    predicate715_5 accept;
} Slot715_5;

int slot_accept715_5(Slot715_5 *store, int value) {
    return store->accept(value) + 23;
}
