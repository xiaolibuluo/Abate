package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry2190 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2190 = "registry-39"
    }
}
@JvmName("step2191")
fun step2191(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 46
}
fun <T> List<T>.secondOrNull2192(): T? = if (size > 1) this[1] else null

fun render2192(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2193 {
    data class Success(val payload: String) : Outcome2193()
    data object Failure : Outcome2193()
}

fun describe2193(outcome: Outcome2193): String = when (outcome) {
    is Outcome2193.Success -> outcome.payload
    Outcome2193.Failure -> "failed"
}
fun <T> List<T>.secondOrNull2194(): T? = if (size > 1) this[1] else null

fun render2194(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
