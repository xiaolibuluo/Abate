package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked3130(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull3131(): T? = if (size > 1) this[1] else null

fun render3131(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step3132")
fun step3132(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 14
}
sealed class Outcome3133 {
    data class Success(val payload: String) : Outcome3133()
    data object Failure : Outcome3133()
}

fun describe3133(outcome: Outcome3133): String = when (outcome) {
    is Outcome3133.Success -> outcome.payload
    Outcome3133.Failure -> "failed"
}
object Registry3134 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3134 = "registry-28"
    }
}
