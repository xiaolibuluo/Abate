package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked12870(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome12871 {
    data class Success(val payload: String) : Outcome12871()
    data object Failure : Outcome12871()
}

fun describe12871(outcome: Outcome12871): String = when (outcome) {
    is Outcome12871.Success -> outcome.payload
    Outcome12871.Failure -> "failed"
}
fun <T, R> List<T>.mapChecked12872(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step12873")
fun step12873(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 73
}
fun <T> List<T>.secondOrNull12874(): T? = if (size > 1) this[1] else null

fun render12874(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
