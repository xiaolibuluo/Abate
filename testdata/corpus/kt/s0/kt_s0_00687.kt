package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5640(): T? = if (size > 1) this[1] else null

fun render5640(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step5641")
fun step5641(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
@JvmName("step5642")
fun step5642(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 76
}
sealed class Outcome5643 {
    data class Success(val payload: String) : Outcome5643()
    data object Failure : Outcome5643()
}

fun describe5643(outcome: Outcome5643): String = when (outcome) {
    is Outcome5643.Success -> outcome.payload
    Outcome5643.Failure -> "failed"
}
@JvmName("step5644")
fun step5644(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
object Registry5645 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5645 = "registry-97"
    }
}
fun <T> List<T>.secondOrNull5646(): T? = if (size > 1) this[1] else null

fun render5646(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
