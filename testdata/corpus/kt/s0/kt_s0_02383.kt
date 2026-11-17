package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull19590(): T? = if (size > 1) this[1] else null

fun render19590(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry19591 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG19591 = "registry-65"
    }
}
sealed class Outcome19592 {
    data class Success(val payload: String) : Outcome19592()
    data object Failure : Outcome19592()
}

fun describe19592(outcome: Outcome19592): String = when (outcome) {
    is Outcome19592.Success -> outcome.payload
    Outcome19592.Failure -> "failed"
}
