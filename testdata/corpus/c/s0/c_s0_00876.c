
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind604_0 { KIND_EMPTY604_0 = 0, KIND_FULL604_0 = 97 };

union Cell604_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind604_0(union Cell604_0 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL604_0 : KIND_EMPTY604_0;
}
enum Kind604_1 { KIND_EMPTY604_1 = 0, KIND_FULL604_1 = 7 };

union Cell604_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind604_1(union Cell604_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL604_1 : KIND_EMPTY604_1;
}
typedef int (*predicate604_2)(int value);

typedef struct Slot604_2 {
    int slots[16];
    predicate604_2 accept;
} Slot604_2;

int slot_accept604_2(Slot604_2 *store, int value) {
    return store->accept(value) + 14;
}
#define MASK604_3(v) (((v) << 4) | ((v) & 0x0F))

int classify604_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK604_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
