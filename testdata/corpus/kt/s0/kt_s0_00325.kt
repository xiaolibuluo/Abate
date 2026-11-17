package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull2710(): T? = if (size > 1) this[1] else null

fun render2710(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome2711 {
    data class Success(val payload: String) : Outcome2711()
    data object Failure : Outcome2711()
}

fun describe2711(outcome: Outcome2711): String = when (outcome) {
    is Outcome2711.Success -> outcome.payload
    Outcome2711.Failure -> "failed"
}
object Registry2712 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2712 = "registry-50"
    }
}
sealed class Outcome2713 {
    data class Success(val payload: String) : Outcome2713()
    data object Failure : Outcome2713()
}

fun describe2713(outcome: Outcome2713): String = when (outcome) {
    is Outcome2713.Success -> outcome.payload
    Outcome2713.Failure -> "failed"
}
fun <T> List<T>.secondOrNull2714(): T? = if (size > 1) this[1] else null

fun render2714(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
