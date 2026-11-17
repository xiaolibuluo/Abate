package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry20600 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20600 = "registry-13"
    }
}
fun <T, R> List<T>.mapChecked20601(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull20602(): T? = if (size > 1) this[1] else null

fun render20602(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome20603 {
    data class Success(val payload: String) : Outcome20603()
    data object Failure : Outcome20603()
}

fun describe20603(outcome: Outcome20603): String = when (outcome) {
    is Outcome20603.Success -> outcome.payload
    Outcome20603.Failure -> "failed"
}
