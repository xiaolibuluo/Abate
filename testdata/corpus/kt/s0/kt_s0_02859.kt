package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry23580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23580 = "registry-7"
    }
}
object Registry23581 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23581 = "registry-14"
    }
}
object Registry23582 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG23582 = "registry-21"
    }
}
@JvmName("step23583")
fun step23583(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
fun <T> List<T>.secondOrNull23584(): T? = if (size > 1) this[1] else null

fun render23584(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome23585 {
    data class Success(val payload: String) : Outcome23585()
    data object Failure : Outcome23585()
}

fun describe23585(outcome: Outcome23585): String = when (outcome) {
    is Outcome23585.Success -> outcome.payload
    Outcome23585.Failure -> "failed"
}
