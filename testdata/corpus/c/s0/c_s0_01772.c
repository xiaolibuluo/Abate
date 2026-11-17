
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1189_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1189_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1189_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1189_1 { KIND_EMPTY1189_1 = 0, KIND_FULL1189_1 = 46 };

union Cell1189_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1189_1(union Cell1189_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1189_1 : KIND_EMPTY1189_1;
}
enum Kind1189_2 { KIND_EMPTY1189_2 = 0, KIND_FULL1189_2 = 53 };

union Cell1189_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1189_2(union Cell1189_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL1189_2 : KIND_EMPTY1189_2;
}
struct Node1189_3 {
    int value;
    struct Node1189_3 *next;
};

int length1189_3(const struct Node1189_3 *head) {
    int count = 0;
    const struct Node1189_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
