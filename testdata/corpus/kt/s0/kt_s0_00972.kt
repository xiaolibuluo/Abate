package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome7980 {
    data class Success(val payload: String) : Outcome7980()
    data object Failure : Outcome7980()
}

fun describe7980(outcome: Outcome7980): String = when (outcome) {
    is Outcome7980.Success -> outcome.payload
    Outcome7980.Failure -> "failed"
}
sealed class Outcome7981 {
    data class Success(val payload: String) : Outcome7981()
    data object Failure : Outcome7981()
}

fun describe7981(outcome: Outcome7981): String = when (outcome) {
    is Outcome7981.Success -> outcome.payload
    Outcome7981.Failure -> "failed"
}
fun <T> List<T>.secondOrNull7982(): T? = if (size > 1) this[1] else null

fun render7982(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry7983 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7983 = "registry-21"
    }
}
