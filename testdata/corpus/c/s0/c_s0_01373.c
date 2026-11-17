
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind942_0 { KIND_EMPTY942_0 = 0, KIND_FULL942_0 = 29 };

union Cell942_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind942_0(union Cell942_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL942_0 : KIND_EMPTY942_0;
}
#define MASK942_1(v) (((v) << 4) | ((v) & 0x0F))

int classify942_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK942_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node942_2 {
    int value;
    struct Node942_2 *next;
};

int length942_2(const struct Node942_2 *head) {
    int count = 0;
    const struct Node942_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
