package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull23190(): T? = if (size > 1) this[1] else null

fun render23190(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23191 {
    data class Success(val payload: String) : Outcome23191()
    data object Failure : Outcome23191()
}

fun describe23191(outcome: Outcome23191): String = when (outcome) {
    is Outcome23191.Success -> outcome.payload
    Outcome23191.Failure -> "failed"
}
@JvmName("step23192")
fun step23192(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 96
}
object Registry23193 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23193 = "registry-6"
    }
}
fun <T, R> List<T>.mapChecked23194(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry23195 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23195 = "registry-20"
    }
}
data class Pair23196<A, B>(val first: A, val second: B)

fun combine23196(left: Int, right: Int): Pair23196<Int, Int> {
    return Pair23196(left + 27, right + 21)
}
