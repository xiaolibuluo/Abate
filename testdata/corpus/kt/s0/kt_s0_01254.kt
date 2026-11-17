package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull10350(): T? = if (size > 1) this[1] else null

fun render10350(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry10351 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10351 = "registry-81"
    }
}
sealed class Outcome10352 {
    data class Success(val payload: String) : Outcome10352()
    data object Failure : Outcome10352()
}

fun describe10352(outcome: Outcome10352): String = when (outcome) {
    is Outcome10352.Success -> outcome.payload
    Outcome10352.Failure -> "failed"
}
