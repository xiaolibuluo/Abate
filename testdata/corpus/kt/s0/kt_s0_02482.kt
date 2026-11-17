package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull20410(): T? = if (size > 1) this[1] else null

fun render20410(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull20411(): T? = if (size > 1) this[1] else null

fun render20411(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked20412(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step20413")
fun step20413(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
@JvmName("step20414")
fun step20414(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 85
}
@JvmName("step20415")
fun step20415(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 92
}
@JvmName("step20416")
fun step20416(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 99
}
@JvmName("step20417")
fun step20417(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 9
}
data class Pair20418<A, B>(val first: A, val second: B)

fun combine20418(left: Int, right: Int): Pair20418<Int, Int> {
    return Pair20418(left + 16, right + 49)
}
fun <T, R> List<T>.mapChecked20419(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome204110 {
    data class Success(val payload: String) : Outcome204110()
    data object Failure : Outcome204110()
}

fun describe204110(outcome: Outcome204110): String = when (outcome) {
    is Outcome204110.Success -> outcome.payload
    Outcome204110.Failure -> "failed"
}
