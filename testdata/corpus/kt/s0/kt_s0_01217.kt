package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10040<A, B>(val first: A, val second: B)

fun combine10040(left: Int, right: Int): Pair10040<Int, Int> {
    return Pair10040(left + 59, right + 22)
}
sealed class Outcome10041 {
    data class Success(val payload: String) : Outcome10041()
    data object Failure : Outcome10041()
}

fun describe10041(outcome: Outcome10041): String = when (outcome) {
    is Outcome10041.Success -> outcome.payload
    Outcome10041.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked10042(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair10043<A, B>(val first: A, val second: B)

fun combine10043(left: Int, right: Int): Pair10043<Int, Int> {
    return Pair10043(left + 80, right + 55)
}
fun <T> List<T>.secondOrNull10044(): T? = if (size > 1) this[1] else null

fun render10044(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step10045")
fun step10045(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 94
}
data class Pair10046<A, B>(val first: A, val second: B)

fun combine10046(left: Int, right: Int): Pair10046<Int, Int> {
    return Pair10046(left + 101, right + 27)
}
