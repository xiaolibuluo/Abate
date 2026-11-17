package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21360 {
    data class Success(val payload: String) : Outcome21360()
    data object Failure : Outcome21360()
}

fun describe21360(outcome: Outcome21360): String = when (outcome) {
    is Outcome21360.Success -> outcome.payload
    Outcome21360.Failure -> "failed"
}
fun <T> List<T>.secondOrNull21361(): T? = if (size > 1) this[1] else null

fun render21361(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair21362<A, B>(val first: A, val second: B)

fun combine21362(left: Int, right: Int): Pair21362<Int, Int> {
    return Pair21362(left + 45, right + 38)
}
data class Pair21363<A, B>(val first: A, val second: B)

fun combine21363(left: Int, right: Int): Pair21363<Int, Int> {
    return Pair21363(left + 52, right + 49)
}
data class Pair21364<A, B>(val first: A, val second: B)

fun combine21364(left: Int, right: Int): Pair21364<Int, Int> {
    return Pair21364(left + 59, right + 60)
}
data class Pair21365<A, B>(val first: A, val second: B)

fun combine21365(left: Int, right: Int): Pair21365<Int, Int> {
    return Pair21365(left + 66, right + 10)
}
fun <T, R> List<T>.mapChecked21366(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked21367(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step21368")
fun step21368(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 87
}
data class Pair21369<A, B>(val first: A, val second: B)

fun combine21369(left: Int, right: Int): Pair21369<Int, Int> {
    return Pair21369(left + 94, right + 54)
}
