package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome12840 {
    data class Success(val payload: String) : Outcome12840()
    data object Failure : Outcome12840()
}

fun describe12840(outcome: Outcome12840): String = when (outcome) {
    is Outcome12840.Success -> outcome.payload
    Outcome12840.Failure -> "failed"
}
object Registry12841 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12841 = "registry-20"
    }
}
object Registry12842 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12842 = "registry-27"
    }
}
fun <T> List<T>.secondOrNull12843(): T? = if (size > 1) this[1] else null

fun render12843(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
