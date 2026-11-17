package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked17310(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step17311")
fun step17311(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 11
}
sealed class Outcome17312 {
    data class Success(val payload: String) : Outcome17312()
    data object Failure : Outcome17312()
}

fun describe17312(outcome: Outcome17312): String = when (outcome) {
    is Outcome17312.Success -> outcome.payload
    Outcome17312.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17313(): T? = if (size > 1) this[1] else null

fun render17313(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome17314 {
    data class Success(val payload: String) : Outcome17314()
    data object Failure : Outcome17314()
}

fun describe17314(outcome: Outcome17314): String = when (outcome) {
    is Outcome17314.Success -> outcome.payload
    Outcome17314.Failure -> "failed"
}
fun <T> List<T>.secondOrNull17315(): T? = if (size > 1) this[1] else null

fun render17315(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
