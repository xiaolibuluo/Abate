package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17210 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17210 = "registry-68"
    }
}
sealed class Outcome17211 {
    data class Success(val payload: String) : Outcome17211()
    data object Failure : Outcome17211()
}

fun describe17211(outcome: Outcome17211): String = when (outcome) {
    is Outcome17211.Success -> outcome.payload
    Outcome17211.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17212(): T? = if (size > 1) this[1] else null

fun render17212(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17213 {
    data class Success(val payload: String) : Outcome17213()
    data object Failure : Outcome17213()
}

fun describe17213(outcome: Outcome17213): String = when (outcome) {
    is Outcome17213.Success -> outcome.payload
    Outcome17213.Failure -> "failed"
}
