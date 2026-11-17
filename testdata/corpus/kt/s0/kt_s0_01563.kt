package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull12930(): T? = if (size > 1) this[1] else null

fun render12930(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair12931<A, B>(val first: A, val second: B)

fun combine12931(left: Int, right: Int): Pair12931<Int, Int> {
    return Pair12931(left + 40, right + 43)
}
fun <T, R> List<T>.mapChecked12932(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12933 {
    data class Success(val payload: String) : Outcome12933()
    data object Failure : Outcome12933()
}

fun describe12933(outcome: Outcome12933): String = when (outcome) {
    is Outcome12933.Success -> outcome.payload
    Outcome12933.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12934(): T? = if (size > 1) this[1] else null

fun render12934(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull12935(): T? = if (size > 1) this[1] else null

fun render12935(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
