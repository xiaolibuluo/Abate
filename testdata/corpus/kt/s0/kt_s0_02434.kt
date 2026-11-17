package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20030(): T? = if (size > 1) this[1] else null

fun render20030(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry20031 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20031 = "registry-55"
    }
}
sealed class Outcome20032 {
    data class Success(val payload: String) : Outcome20032()
    data object Failure : Outcome20032()
}

fun describe20032(outcome: Outcome20032): String = when (outcome) {
    is Outcome20032.Success -> outcome.payload
    Outcome20032.Failure -> "failed"
}
