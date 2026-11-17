package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair6500<A, B>(val first: A, val second: B)

fun combine6500(left: Int, right: Int): Pair6500<Int, Int> {
    return Pair6500(left + 16, right + 45)
}
@JvmName("step6501")
fun step6501(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 23
}
data class Pair6502<A, B>(val first: A, val second: B)

fun combine6502(left: Int, right: Int): Pair6502<Int, Int> {
    return Pair6502(left + 30, right + 67)
}
data class Pair6503<A, B>(val first: A, val second: B)

fun combine6503(left: Int, right: Int): Pair6503<Int, Int> {
    return Pair6503(left + 37, right + 17)
}
fun <T> List<T>.secondOrNull6504(): T? = if (size > 1) this[1] else null

fun render6504(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
