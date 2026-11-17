package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step19790")
fun step19790(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
data class Pair19791<A, B>(val first: A, val second: B)

fun combine19791(left: Int, right: Int): Pair19791<Int, Int> {
    return Pair19791(left + 34, right + 26)
}
data class Pair19792<A, B>(val first: A, val second: B)

fun combine19792(left: Int, right: Int): Pair19792<Int, Int> {
    return Pair19792(left + 41, right + 37)
}
fun <T, R> List<T>.mapChecked19793(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome19794 {
    data class Success(val payload: String) : Outcome19794()
    data object Failure : Outcome19794()
}

fun describe19794(outcome: Outcome19794): String = when (outcome) {
    is Outcome19794.Success -> outcome.payload
    Outcome19794.Failure -> "failed"
}
fun <T> List<T>.secondOrNull19795(): T? = if (size > 1) this[1] else null

fun render19795(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair19796<A, B>(val first: A, val second: B)

fun combine19796(left: Int, right: Int): Pair19796<Int, Int> {
    return Pair19796(left + 69, right + 20)
}
