package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry15820 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15820 = "registry-7"
    }
}
fun <T> List<T>.secondOrNull15821(): T? = if (size > 1) this[1] else null

fun render15821(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome15822 {
    data class Success(val payload: String) : Outcome15822()
    data object Failure : Outcome15822()
}

fun describe15822(outcome: Outcome15822): String = when (outcome) {
    is Outcome15822.Success -> outcome.payload
    Outcome15822.Failure -> "failed"
}
object Registry15823 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15823 = "registry-28"
    }
}
sealed class Outcome15824 {
    data class Success(val payload: String) : Outcome15824()
    data object Failure : Outcome15824()
}

fun describe15824(outcome: Outcome15824): String = when (outcome) {
    is Outcome15824.Success -> outcome.payload
    Outcome15824.Failure -> "failed"
}
