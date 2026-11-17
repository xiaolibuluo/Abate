
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1893_0 { KIND_EMPTY1893_0 = 0, KIND_FULL1893_0 = 73 };

union Cell1893_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1893_0(union Cell1893_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1893_0 : KIND_EMPTY1893_0;
}
#define MASK1893_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1893_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1893_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1893_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1893_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1893_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1893_3 { KIND_EMPTY1893_3 = 0, KIND_FULL1893_3 = 94 };

union Cell1893_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1893_3(union Cell1893_3 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1893_3 : KIND_EMPTY1893_3;
}
struct Node1893_4 {
    int value;
    struct Node1893_4 *next;
};

int length1893_4(const struct Node1893_4 *head) {
    int count = 0;
    const struct Node1893_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
enum Kind1893_5 { KIND_EMPTY1893_5 = 0, KIND_FULL1893_5 = 11 };

union Cell1893_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1893_5(union Cell1893_5 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1893_5 : KIND_EMPTY1893_5;
}
typedef int (*predicate1893_6)(int value);

typedef struct Slot1893_6 {
    int slots[16];
    predicate1893_6 accept;
} Slot1893_6;

int slot_accept1893_6(Slot1893_6 *store, int value) {
    return store->accept(value) + 18;
}
