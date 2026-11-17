
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind424_0 { KIND_EMPTY424_0 = 0, KIND_FULL424_0 = 85 };

union Cell424_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind424_0(union Cell424_0 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL424_0 : KIND_EMPTY424_0;
}
#define MASK424_1(v) (((v) << 4) | ((v) & 0x0F))

int classify424_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK424_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind424_2 { KIND_EMPTY424_2 = 0, KIND_FULL424_2 = 99 };

union Cell424_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind424_2(union Cell424_2 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL424_2 : KIND_EMPTY424_2;
}
typedef int (*predicate424_3)(int value);

typedef struct Slot424_3 {
    int slots[16];
    predicate424_3 accept;
} Slot424_3;

int slot_accept424_3(Slot424_3 *store, int value) {
    return store->accept(value) + 9;
}
#define MASK424_4(v) (((v) << 4) | ((v) & 0x0F))

int classify424_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK424_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind424_5 { KIND_EMPTY424_5 = 0, KIND_FULL424_5 = 23 };

union Cell424_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind424_5(union Cell424_5 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL424_5 : KIND_EMPTY424_5;
}
