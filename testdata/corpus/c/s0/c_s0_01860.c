
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1248_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1248_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1248_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1248_1)(int value);

typedef struct Slot1248_1 {
    int slots[16];
    predicate1248_1 accept;
} Slot1248_1;

int slot_accept1248_1(Slot1248_1 *store, int value) {
    return store->accept(value) + 37;
}
enum Kind1248_2 { KIND_EMPTY1248_2 = 0, KIND_FULL1248_2 = 44 };

union Cell1248_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1248_2(union Cell1248_2 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1248_2 : KIND_EMPTY1248_2;
}
#if defined(__STDC_VERSION__)
int variant1248_3(int value, ...) {
    return value + 51;
}
#else
int variant1248_3(int value) {
    return value + 51;
}
#endif
#if defined(__STDC_VERSION__)
int variant1248_4(int value, ...) {
    return value + 58;
}
#else
int variant1248_4(int value) {
    return value + 58;
}
#endif
typedef int (*predicate1248_5)(int value);

typedef struct Slot1248_5 {
    int slots[16];
    predicate1248_5 accept;
} Slot1248_5;

int slot_accept1248_5(Slot1248_5 *store, int value) {
    return store->accept(value) + 65;
}
struct Node1248_6 {
    int value;
    struct Node1248_6 *next;
};

int length1248_6(const struct Node1248_6 *head) {
    int count = 0;
    const struct Node1248_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
