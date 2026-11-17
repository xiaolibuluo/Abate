package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull22100(): T? = if (size > 1) this[1] else null

fun render22100(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry22101 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22101 = "registry-30"
    }
}
sealed class Outcome22102 {
    data class Success(val payload: String) : Outcome22102()
    data object Failure : Outcome22102()
}

fun describe22102(outcome: Outcome22102): String = when (outcome) {
    is Outcome22102.Success -> outcome.payload
    Outcome22102.Failure -> "failed"
}
object Registry22103 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22103 = "registry-44"
    }
}
