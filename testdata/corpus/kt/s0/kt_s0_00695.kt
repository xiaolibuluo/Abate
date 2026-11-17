package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair5700<A, B>(val first: A, val second: B)

fun combine5700(left: Int, right: Int): Pair5700<Int, Int> {
    return Pair5700(left + 43, right + 34)
}
@JvmName("step5701")
fun step5701(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 50
}
data class Pair5702<A, B>(val first: A, val second: B)

fun combine5702(left: Int, right: Int): Pair5702<Int, Int> {
    return Pair5702(left + 57, right + 56)
}
fun <T, R> List<T>.mapChecked5703(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5704")
fun step5704(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
data class Pair5705<A, B>(val first: A, val second: B)

fun combine5705(left: Int, right: Int): Pair5705<Int, Int> {
    return Pair5705(left + 78, right + 28)
}
fun <T> List<T>.secondOrNull5706(): T? = if (size > 1) this[1] else null

fun render5706(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5707 {
    data class Success(val payload: String) : Outcome5707()
    data object Failure : Outcome5707()
}

fun describe5707(outcome: Outcome5707): String = when (outcome) {
    is Outcome5707.Success -> outcome.payload
    Outcome5707.Failure -> "failed"
}
