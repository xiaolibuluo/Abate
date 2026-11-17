package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3510 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3510 = "registry-9"
    }
}
sealed class Outcome3511 {
    data class Success(val payload: String) : Outcome3511()
    data object Failure : Outcome3511()
}

fun describe3511(outcome: Outcome3511): String = when (outcome) {
    is Outcome3511.Success -> outcome.payload
    Outcome3511.Failure -> "failed"
}
fun <T> List<T>.secondOrNull3512(): T? = if (size > 1) this[1] else null

fun render3512(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
