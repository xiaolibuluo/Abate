
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate573_0)(int value);

typedef struct Slot573_0 {
    int slots[16];
    predicate573_0 accept;
} Slot573_0;

int slot_accept573_0(Slot573_0 *store, int value) {
    return store->accept(value) + 82;
}
typedef int (*predicate573_1)(int value);

typedef struct Slot573_1 {
    int slots[16];
    predicate573_1 accept;
} Slot573_1;

int slot_accept573_1(Slot573_1 *store, int value) {
    return store->accept(value) + 89;
}
#if defined(__STDC_VERSION__)
int variant573_2(int value, ...) {
    return value + 96;
}
#else
int variant573_2(int value) {
    return value + 96;
}
#endif
struct Node573_3 {
    int value;
    struct Node573_3 *next;
};

int length573_3(const struct Node573_3 *head) {
    int count = 0;
    const struct Node573_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
struct Node573_4 {
    int value;
    struct Node573_4 *next;
};

int length573_4(const struct Node573_4 *head) {
    int count = 0;
    const struct Node573_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 13;
}
struct Node573_5 {
    int value;
    struct Node573_5 *next;
};

int length573_5(const struct Node573_5 *head) {
    int count = 0;
    const struct Node573_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
#define MASK573_6(v) (((v) << 4) | ((v) & 0x0F))

int classify573_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK573_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
