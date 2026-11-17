package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23140 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23140 = "registry-17"
    }
}
fun <T> List<T>.secondOrNull23141(): T? = if (size > 1) this[1] else null

fun render23141(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair23142<A, B>(val first: A, val second: B)

fun combine23142(left: Int, right: Int): Pair23142<Int, Int> {
    return Pair23142(left + 31, right + 64)
}
fun <T, R> List<T>.mapChecked23143(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome23144 {
    data class Success(val payload: String) : Outcome23144()
    data object Failure : Outcome23144()
}

fun describe23144(outcome: Outcome23144): String = when (outcome) {
    is Outcome23144.Success -> outcome.payload
    Outcome23144.Failure -> "failed"
}
@JvmName("step23145")
fun step23145(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 52
}
