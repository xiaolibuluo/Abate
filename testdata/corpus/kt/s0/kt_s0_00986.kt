package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

class Counter8080 {
    var current: Int by Delegates.observable(33) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
@JvmName("step8081")
fun step8081(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 40
}
@JvmName("step8082")
fun step8082(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 47
}
class Counter8083 {
    var current: Int by Delegates.observable(54) { _, _, new -> println(new) }

    suspend fun advance(): Int {
        try {
            current += 1
        } catch (failure: IllegalStateException) {
            current = 0
        } finally {
            println("advanced")
        }
        return current
    }
}
fun <T, R> List<T>.mapChecked8084(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair8085<A, B>(val first: A, val second: B)

fun combine8085(left: Int, right: Int): Pair8085<Int, Int> {
    return Pair8085(left + 68, right + 47)
}
