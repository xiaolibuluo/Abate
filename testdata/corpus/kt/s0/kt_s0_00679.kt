package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull5570(): T? = if (size > 1) this[1] else null

fun render5570(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull5571(): T? = if (size > 1) this[1] else null

fun render5571(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome5572 {
    data class Success(val payload: String) : Outcome5572()
    data object Failure : Outcome5572()
}

fun describe5572(outcome: Outcome5572): String = when (outcome) {
    is Outcome5572.Success -> outcome.payload
    Outcome5572.Failure -> "failed"
}
object Registry5573 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5573 = "registry-89"
    }
}
sealed class Outcome5574 {
    data class Success(val payload: String) : Outcome5574()
    data object Failure : Outcome5574()
}

fun describe5574(outcome: Outcome5574): String = when (outcome) {
    is Outcome5574.Success -> outcome.payload
    Outcome5574.Failure -> "failed"
}
