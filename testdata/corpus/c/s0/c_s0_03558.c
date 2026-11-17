
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2402_0 { KIND_EMPTY2402_0 = 0, KIND_FULL2402_0 = 94 };

union Cell2402_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2402_0(union Cell2402_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2402_0 : KIND_EMPTY2402_0;
}
enum Kind2402_1 { KIND_EMPTY2402_1 = 0, KIND_FULL2402_1 = 101 };

union Cell2402_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2402_1(union Cell2402_1 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL2402_1 : KIND_EMPTY2402_1;
}
struct Node2402_2 {
    int value;
    struct Node2402_2 *next;
};

int length2402_2(const struct Node2402_2 *head) {
    int count = 0;
    const struct Node2402_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
