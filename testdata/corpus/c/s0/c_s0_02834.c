
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1924_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1924_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1924_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1924_1 { KIND_EMPTY1924_1 = 0, KIND_FULL1924_1 = 95 };

union Cell1924_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1924_1(union Cell1924_1 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1924_1 : KIND_EMPTY1924_1;
}
typedef int (*predicate1924_2)(int value);

typedef struct Slot1924_2 {
    int slots[16];
    predicate1924_2 accept;
} Slot1924_2;

int slot_accept1924_2(Slot1924_2 *store, int value) {
    return store->accept(value) + 5;
}
enum Kind1924_3 { KIND_EMPTY1924_3 = 0, KIND_FULL1924_3 = 12 };

union Cell1924_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1924_3(union Cell1924_3 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1924_3 : KIND_EMPTY1924_3;
}
#define MASK1924_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1924_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1924_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
