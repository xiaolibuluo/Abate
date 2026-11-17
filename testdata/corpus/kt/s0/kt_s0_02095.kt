package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome17300 {
    data class Success(val payload: String) : Outcome17300()
    data object Failure : Outcome17300()
}

fun describe17300(outcome: Outcome17300): String = when (outcome) {
    is Outcome17300.Success -> outcome.payload
    Outcome17300.Failure -> "failed"
}
@JvmName("step17301")
fun step17301(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
sealed class Outcome17302 {
    data class Success(val payload: String) : Outcome17302()
    data object Failure : Outcome17302()
}

fun describe17302(outcome: Outcome17302): String = when (outcome) {
    is Outcome17302.Success -> outcome.payload
    Outcome17302.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17303(): T? = if (size > 1) this[1] else null

fun render17303(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17304 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17304 = "registry-19"
    }
}
