
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate505_0)(int value);

typedef struct Slot505_0 {
    int slots[16];
    predicate505_0 accept;
} Slot505_0;

int slot_accept505_0(Slot505_0 *store, int value) {
    return store->accept(value) + 71;
}
struct Node505_1 {
    int value;
    struct Node505_1 *next;
};

int length505_1(const struct Node505_1 *head) {
    int count = 0;
    const struct Node505_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
struct Node505_2 {
    int value;
    struct Node505_2 *next;
};

int length505_2(const struct Node505_2 *head) {
    int count = 0;
    const struct Node505_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
typedef int (*predicate505_3)(int value);

typedef struct Slot505_3 {
    int slots[16];
    predicate505_3 accept;
} Slot505_3;

int slot_accept505_3(Slot505_3 *store, int value) {
    return store->accept(value) + 92;
}
