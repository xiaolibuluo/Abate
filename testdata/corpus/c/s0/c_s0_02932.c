
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1982_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1982_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1982_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1982_1)(int value);

typedef struct Slot1982_1 {
    int slots[16];
    predicate1982_1 accept;
} Slot1982_1;

int slot_accept1982_1(Slot1982_1 *store, int value) {
    return store->accept(value) + 73;
}
enum Kind1982_2 { KIND_EMPTY1982_2 = 0, KIND_FULL1982_2 = 80 };

union Cell1982_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1982_2(union Cell1982_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1982_2 : KIND_EMPTY1982_2;
}
struct Node1982_3 {
    int value;
    struct Node1982_3 *next;
};

int length1982_3(const struct Node1982_3 *head) {
    int count = 0;
    const struct Node1982_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
typedef int (*predicate1982_4)(int value);

typedef struct Slot1982_4 {
    int slots[16];
    predicate1982_4 accept;
} Slot1982_4;

int slot_accept1982_4(Slot1982_4 *store, int value) {
    return store->accept(value) + 94;
}
enum Kind1982_5 { KIND_EMPTY1982_5 = 0, KIND_FULL1982_5 = 101 };

union Cell1982_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1982_5(union Cell1982_5 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1982_5 : KIND_EMPTY1982_5;
}
