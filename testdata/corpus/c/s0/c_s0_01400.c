
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind961_0 { KIND_EMPTY961_0 = 0, KIND_FULL961_0 = 82 };

union Cell961_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind961_0(union Cell961_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL961_0 : KIND_EMPTY961_0;
}
enum Kind961_1 { KIND_EMPTY961_1 = 0, KIND_FULL961_1 = 89 };

union Cell961_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind961_1(union Cell961_1 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL961_1 : KIND_EMPTY961_1;
}
typedef int (*predicate961_2)(int value);

typedef struct Slot961_2 {
    int slots[16];
    predicate961_2 accept;
} Slot961_2;

int slot_accept961_2(Slot961_2 *store, int value) {
    return store->accept(value) + 96;
}
typedef int (*predicate961_3)(int value);

typedef struct Slot961_3 {
    int slots[16];
    predicate961_3 accept;
} Slot961_3;

int slot_accept961_3(Slot961_3 *store, int value) {
    return store->accept(value) + 6;
}
enum Kind961_4 { KIND_EMPTY961_4 = 0, KIND_FULL961_4 = 13 };

union Cell961_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind961_4(union Cell961_4 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL961_4 : KIND_EMPTY961_4;
}
#define MASK961_5(v) (((v) << 4) | ((v) & 0x0F))

int classify961_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK961_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
