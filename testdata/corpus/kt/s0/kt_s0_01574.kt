package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome13040 {
    data class Success(val payload: String) : Outcome13040()
    data object Failure : Outcome13040()
}

fun describe13040(outcome: Outcome13040): String = when (outcome) {
    is Outcome13040.Success -> outcome.payload
    Outcome13040.Failure -> "failed"
}
fun <T> List<T>.secondOrNull13041(): T? = if (size > 1) this[1] else null

fun render13041(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair13042<A, B>(val first: A, val second: B)

fun combine13042(left: Int, right: Int): Pair13042<Int, Int> {
    return Pair13042(left + 93, right + 9)
}
data class Pair13043<A, B>(val first: A, val second: B)

fun combine13043(left: Int, right: Int): Pair13043<Int, Int> {
    return Pair13043(left + 100, right + 20)
}
data class Pair13044<A, B>(val first: A, val second: B)

fun combine13044(left: Int, right: Int): Pair13044<Int, Int> {
    return Pair13044(left + 10, right + 31)
}
data class Pair13045<A, B>(val first: A, val second: B)

fun combine13045(left: Int, right: Int): Pair13045<Int, Int> {
    return Pair13045(left + 17, right + 42)
}
@JvmName("step13046")
fun step13046(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
@JvmName("step13047")
fun step13047(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 31
}
fun <T, R> List<T>.mapChecked13048(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
