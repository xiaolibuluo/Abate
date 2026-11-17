package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6760")
fun step6760(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 63
}
object Registry6761 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6761 = "registry-70"
    }
}
fun <T> List<T>.secondOrNull6762(): T? = if (size > 1) this[1] else null

fun render6762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked6763(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome6764 {
    data class Success(val payload: String) : Outcome6764()
    data object Failure : Outcome6764()
}

fun describe6764(outcome: Outcome6764): String = when (outcome) {
    is Outcome6764.Success -> outcome.payload
    Outcome6764.Failure -> "failed"
}
data class Pair6765<A, B>(val first: A, val second: B)

fun combine6765(left: Int, right: Int): Pair6765<Int, Int> {
    return Pair6765(left + 98, right + 38)
}
fun <T, R> List<T>.mapChecked6766(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step6767")
fun step6767(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
sealed class Outcome6768 {
    data class Success(val payload: String) : Outcome6768()
    data object Failure : Outcome6768()
}

fun describe6768(outcome: Outcome6768): String = when (outcome) {
    is Outcome6768.Success -> outcome.payload
    Outcome6768.Failure -> "failed"
}
