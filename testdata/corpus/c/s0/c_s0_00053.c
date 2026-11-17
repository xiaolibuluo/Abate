
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK33_0(v) (((v) << 4) | ((v) & 0x0F))

int classify33_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK33_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node33_1 {
    int value;
    struct Node33_1 *next;
};

int length33_1(const struct Node33_1 *head) {
    int count = 0;
    const struct Node33_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 53;
}
enum Kind33_2 { KIND_EMPTY33_2 = 0, KIND_FULL33_2 = 60 };

union Cell33_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind33_2(union Cell33_2 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL33_2 : KIND_EMPTY33_2;
}
enum Kind33_3 { KIND_EMPTY33_3 = 0, KIND_FULL33_3 = 67 };

union Cell33_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind33_3(union Cell33_3 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL33_3 : KIND_EMPTY33_3;
}
#define MASK33_4(v) (((v) << 4) | ((v) & 0x0F))

int classify33_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK33_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate33_5)(int value);

typedef struct Slot33_5 {
    int slots[16];
    predicate33_5 accept;
} Slot33_5;

int slot_accept33_5(Slot33_5 *store, int value) {
    return store->accept(value) + 81;
}
