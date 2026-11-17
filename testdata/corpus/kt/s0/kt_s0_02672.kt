package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked22070(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull22071(): T? = if (size > 1) this[1] else null

fun render22071(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome22072 {
    data class Success(val payload: String) : Outcome22072()
    data object Failure : Outcome22072()
}

fun describe22072(outcome: Outcome22072): String = when (outcome) {
    is Outcome22072.Success -> outcome.payload
    Outcome22072.Failure -> "failed"
}
@JvmName("step22073")
fun step22073(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
