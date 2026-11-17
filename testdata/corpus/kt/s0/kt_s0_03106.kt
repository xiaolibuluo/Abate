package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25500 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25500 = "registry-78"
    }
}
sealed class Outcome25501 {
    data class Success(val payload: String) : Outcome25501()
    data object Failure : Outcome25501()
}

fun describe25501(outcome: Outcome25501): String = when (outcome) {
    is Outcome25501.Success -> outcome.payload
    Outcome25501.Failure -> "failed"
}
fun <T> List<T>.secondOrNull25502(): T? = if (size > 1) this[1] else null

fun render25502(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome25503 {
    data class Success(val payload: String) : Outcome25503()
    data object Failure : Outcome25503()
}

fun describe25503(outcome: Outcome25503): String = when (outcome) {
    is Outcome25503.Success -> outcome.payload
    Outcome25503.Failure -> "failed"
}
