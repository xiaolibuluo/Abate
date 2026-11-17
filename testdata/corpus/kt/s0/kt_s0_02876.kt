package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23730 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23730 = "registry-8"
    }
}
fun <T> List<T>.secondOrNull23731(): T? = if (size > 1) this[1] else null

fun render23731(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step23732")
fun step23732(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
sealed class Outcome23733 {
    data class Success(val payload: String) : Outcome23733()
    data object Failure : Outcome23733()
}

fun describe23733(outcome: Outcome23733): String = when (outcome) {
    is Outcome23733.Success -> outcome.payload
    Outcome23733.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked23734(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step23735")
fun step23735(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 43
}
sealed class Outcome23736 {
    data class Success(val payload: String) : Outcome23736()
    data object Failure : Outcome23736()
}

fun describe23736(outcome: Outcome23736): String = when (outcome) {
    is Outcome23736.Success -> outcome.payload
    Outcome23736.Failure -> "failed"
}
data class Pair23737<A, B>(val first: A, val second: B)

fun combine23737(left: Int, right: Int): Pair23737<Int, Int> {
    return Pair23737(left + 57, right + 44)
}
