
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node654_0 {
    int value;
    struct Node654_0 *next;
};

int length654_0(const struct Node654_0 *head) {
    int count = 0;
    const struct Node654_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
#define MASK654_1(v) (((v) << 4) | ((v) & 0x0F))

int classify654_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK654_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate654_2)(int value);

typedef struct Slot654_2 {
    int slots[16];
    predicate654_2 accept;
} Slot654_2;

int slot_accept654_2(Slot654_2 *store, int value) {
    return store->accept(value) + 82;
}
typedef int (*predicate654_3)(int value);

typedef struct Slot654_3 {
    int slots[16];
    predicate654_3 accept;
} Slot654_3;

int slot_accept654_3(Slot654_3 *store, int value) {
    return store->accept(value) + 89;
}
