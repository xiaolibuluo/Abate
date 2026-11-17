package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair12110<A, B>(val first: A, val second: B)

fun combine12110(left: Int, right: Int): Pair12110<Int, Int> {
    return Pair12110(left + 34, right + 68)
}
data class Pair12111<A, B>(val first: A, val second: B)

fun combine12111(left: Int, right: Int): Pair12111<Int, Int> {
    return Pair12111(left + 41, right + 18)
}
object Registry12112 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12112 = "registry-48"
    }
}
fun <T, R> List<T>.mapChecked12113(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull12114(): T? = if (size > 1) this[1] else null

fun render12114(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12115")
fun step12115(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 69
}
sealed class Outcome12116 {
    data class Success(val payload: String) : Outcome12116()
    data object Failure : Outcome12116()
}

fun describe12116(outcome: Outcome12116): String = when (outcome) {
    is Outcome12116.Success -> outcome.payload
    Outcome12116.Failure -> "failed"
}
fun <T> List<T>.secondOrNull12117(): T? = if (size > 1) this[1] else null

fun render12117(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step12118")
fun step12118(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 90
}
