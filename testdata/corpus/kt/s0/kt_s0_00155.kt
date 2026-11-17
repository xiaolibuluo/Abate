package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair1260<A, B>(val first: A, val second: B)

fun combine1260(left: Int, right: Int): Pair1260<Int, Int> {
    return Pair1260(left + 91, right + 37)
}
fun <T, R> List<T>.mapChecked1261(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull1262(): T? = if (size > 1) this[1] else null

fun render1262(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome1263 {
    data class Success(val payload: String) : Outcome1263()
    data object Failure : Outcome1263()
}

fun describe1263(outcome: Outcome1263): String = when (outcome) {
    is Outcome1263.Success -> outcome.payload
    Outcome1263.Failure -> "failed"
}
data class Pair1264<A, B>(val first: A, val second: B)

fun combine1264(left: Int, right: Int): Pair1264<Int, Int> {
    return Pair1264(left + 22, right + 20)
}
data class Pair1265<A, B>(val first: A, val second: B)

fun combine1265(left: Int, right: Int): Pair1265<Int, Int> {
    return Pair1265(left + 29, right + 31)
}
