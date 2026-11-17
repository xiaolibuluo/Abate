
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK698_0(v) (((v) << 4) | ((v) & 0x0F))

int classify698_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK698_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate698_1)(int value);

typedef struct Slot698_1 {
    int slots[16];
    predicate698_1 accept;
} Slot698_1;

int slot_accept698_1(Slot698_1 *store, int value) {
    return store->accept(value) + 65;
}
#define MASK698_2(v) (((v) << 4) | ((v) & 0x0F))

int classify698_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK698_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind698_3 { KIND_EMPTY698_3 = 0, KIND_FULL698_3 = 79 };

union Cell698_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind698_3(union Cell698_3 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL698_3 : KIND_EMPTY698_3;
}
